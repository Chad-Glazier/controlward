import styles from "./Header.module.css"
import logo from "../../assets/logo.png"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import { Link } from "react-router";
import { useRef } from "react";

type HeaderProps = {
    onSearch: SearchProps["onSearch"]
}

function Header({ onSearch }: HeaderProps) {



    return <header className={styles.header} id="header">
        <Link to="/">
            <div className={styles.brand}>
                <img className={styles.logo} src={logo} alt="logo" />
                <h1 className={styles.title}>Control Ward</h1>
            </div>        
        </Link>
        <Search onSearch={onSearch} />
        <div className={styles.actions}>
            <button className={styles.button}>b</button>
        </div>
    </header>
}

type SearchProps = {
    onSearch: (gameName: string, tagLine: string) => Promise<void>
}

function Search({ onSearch }: SearchProps ) {

    const inputRef = useRef<HTMLInputElement | null>(null)

    const handler = useRef((ev: KeyboardEvent) => {
        if (ev.key !== "Enter") {
            return
        }

        inputRef.current!.value
    })

    return (
        <div className={styles.search}>
            <input
                ref={inputRef}
                className={styles.searchInput}
                type="text"
                placeholder="Hide on bush#KR1"
                onFocus={() => {
                    document.addEventListener("keypress", handler.current)
                }}
                onBlur={() => {
                    document.removeEventListener("keypress", handler.current)
                }}
            />
            <button
                className={styles.searchButton}
            >
                <FontAwesomeIcon icon={faSearch} />
            </button>
        </div>
    );
}

export default Header 
