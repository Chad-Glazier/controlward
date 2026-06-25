import styles from "./Header.module.css"
import logo from "../../assets/logo.png"
import { Link } from "react-router";

type HeaderProps = {
    hidden?: boolean
}

function Header({ hidden }: HeaderProps) {
    return <header 
        className={styles.header + " " + (hidden ? styles.hidden : "")} 
        id="header"
    >
        <Link to="/">
            <div className={styles.brand}>
                <img className={styles.logo} src={logo} alt="logo" />
                <h1 className={styles.title}>Control Ward</h1>
            </div>        
        </Link>
        <div className={styles.actions}>
            <button className={styles.button}>b</button>
        </div>
    </header>
}

export default Header 
