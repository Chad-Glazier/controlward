import { useEffect, useState, type PropsWithChildren } from "react"
import styles from "./Layout.module.css"
import Header from "./Header"
import Footer from "./Footer"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faArrowUp } from "@fortawesome/free-solid-svg-icons"

type LayoutProps = PropsWithChildren<{}>

function Layout({ children }: LayoutProps) {

    const [showScrollupButton, setShowScrollupButton] = useState(false)

    useEffect(() => {
        window.addEventListener("scroll", _ => {
            if (window.scrollY > 300) {
                setShowScrollupButton(true)
            } else {
                setShowScrollupButton(false)
            }
        })
    }, [])

    return <div className={styles.container}>
        <Header onSearch={async () => {}} />
        <main className={styles.main}>
            {children}
        </main>
        <ScrollupButton show={showScrollupButton} />
        <Footer />
    </div>
}

function ScrollupButton({ show }: { show: boolean }) {
    return <button 
        className={styles.scrollup + " " + (show ? "" : styles.hidden)}
        onClick={() => {
            window.scrollTo(0, 0)
        }}
    >
        <FontAwesomeIcon icon={faArrowUp} />
    </button>  
}

export default Layout
