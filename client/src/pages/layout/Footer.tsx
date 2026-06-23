import styles from "./Footer.module.css";

function Footer() {
    const year = new Date().getFullYear();

    return (
        <footer className={styles.footer}>
            <div className={styles.content}>
                <small>&#xA9; {year} Control Ward. All rights reserved.</small>
                <small>Not affiliated with Riot Games</small>
                <div className={styles.contact}>
                    <small>
                        <span>Contact:</span>&nbsp;
                        <a href="mailto:chadglazier@outlook.com">
                            chadglazier@outlook.com
                        </a>                        
                    </small>
                </div>
            </div>
        </footer>
    );
}

export default Footer;