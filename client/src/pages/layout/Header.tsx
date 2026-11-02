import styles from "./Header.module.css"
import Logo from "../../components/Logo";
import PlayerSearch from "../../components/PlayerSearch";
import { useNavigate } from "react-router";

type HeaderProps = {
    hidden?: boolean
}

function Header({ hidden }: HeaderProps) {
    const navigate = useNavigate()

    return <header 
        className={styles.header + " " + (hidden ? styles.hidden : "")} 
        id="header"
    >
        <Logo 
            height="40px"
            clickable
        />
        <PlayerSearch 
            onSearch={async (puuid) => {
                window.scrollTo(0, 0)
                navigate(`/profile/${puuid}`)
            }}
            className={styles.search}
        />
    </header>
}

export default Header 
