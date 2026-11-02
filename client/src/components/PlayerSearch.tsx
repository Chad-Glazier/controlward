import { useRef, useState } from "react"
import styles from "./PlayerSearch.module.css"
import { validateRiotAccount } from "../cw/validators"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faSearch, faSpinner } from '@fortawesome/free-solid-svg-icons';
import getPuuid from "../cw/getPuuid";

type PlayerSearchProps = {
    onSearch: (puuid: string) => Promise<void>
    className?: string
}

export default function PlayerSearch({
    onSearch, className,
}: PlayerSearchProps) {
    const [loading, setLoading] = useState(false)
    const [focused, setFocused] = useState(false)
    const [showError, setShowError] = useState(false)
    const [error, setError] = useState("")
    const input = useRef<HTMLInputElement | null>(null)
    const errTimeout = useRef<number>(0)

    const handler = useRef(async (ev: KeyboardEvent) => {
        if (ev.key !== "Enter") {
            return
        }

        const [gameName, tagLine, err] = validateRiotAccount(input.current!.value)
        if (err != null) {
            setError(err)
            setShowError(true)
            return
        }

        setShowError(false)
        setError("")
        
        setLoading(true)
        let puuid = await getPuuid(gameName, tagLine)
        if (puuid === null) {
            setError("User not found")
            setShowError(true)
        } else {
            await onSearch(puuid)
        }
        setLoading(false)
    })

    return <div className={styles.container + " " + (className ?? "")}>
        <search className={styles.search}>
            <div className={styles.icon}>
                {loading 
                    ? <FontAwesomeIcon key="a" icon={faSpinner} spin />
                    : <FontAwesomeIcon key="b" icon={faSearch} />
                }
            </div>
            <input 
                ref={input}
                className={styles.input}
                type="text"
                placeholder="Hide on bush#KR1"
                onFocus={() => {
                    setFocused(true)
                    document.addEventListener("keypress", handler.current)
                }}
                onBlur={() => {
                    setFocused(false)
                    setShowError(false)
                    document.removeEventListener("keypress", handler.current)
                }}
                onChange={ev => {
                    setShowError(false)
                    setError("")
                    clearTimeout(errTimeout.current)
                    errTimeout.current = setTimeout(() => {
                        if (ev.target.value === "") {
                            return
                        }
                        const [_, __, err] = validateRiotAccount(ev.target.value)
                        if (err !== null) {
                            setError(err)
                        }
                    }, 400)
                }}
            >
            </input>
        </search>    
        <div 
            className={styles.error 
                + " " + (error === "" 
                    ? styles.valid
                    : styles.invalid)
                + " " + (showError
                    ? styles.show
                    : styles.hide)
                + " " + (focused 
                    ? styles.focused
                    : styles.blurred)            
            }>
            {(showError && error) || <>&nbsp;</>}
        </div>
    </div> 
}
