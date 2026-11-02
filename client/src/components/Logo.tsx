import { Link } from "react-router"
import styles from "./Logo.module.css"

type LogoProps = {
    height: string
    clickable?: boolean
}

function Logo({ height, clickable }: LogoProps) {
    if (clickable) {
        return <Link to="/">
            <div className={styles.container}>
                <h1 
                    className={styles.header + " " + styles.clickable}
                    style={{ 
                        fontSize: `calc(${height} / 2)`,
                        borderWidth: `calc(${height} / 16)`, 
                        paddingRight: `calc(${height} / 4)`,
                    }}    
                >
                    Control<br />Ward
                </h1>
            </div>
        </Link>
    } else {
        return <div className={styles.container}>
            <div className={styles.container}>
                <h1 
                    className={styles.header}
                    style={{ 
                        fontSize: `calc(${height} / 2)`,
                        borderWidth: `calc(${height} / 16)`, 
                        paddingRight: `calc(${height} / 4)`,
                    }}    
                >
                    Control<br />Ward
                </h1>
            </div>
        </div>
    }
}

export default Logo
