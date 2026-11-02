import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import styles from "./Footer.module.css";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

function Footer() {
    return (
        <footer className={styles.footer}>
            <div className={styles.content}>
                <ul>
                    <li>
                        <FontAwesomeIcon icon={faGithub} />
                        &nbsp;
                        <a 
                            target="_blank"
                            href="https://github.com/Chad-Glazier/controlward"
                        >
                            Source
                        </a>
                    </li>
                </ul>
                <ul>
                    <li>
                        Control Ward is not endorsed by Riot Games and does not 
                        reflect the views or opinions of Riot Games or anyone 
                        officially involved in producing or managing League of 
                        Legends. League of Legends and Riot Games are trademarks 
                        or registered trademarks of Riot Games, Inc. League of 
                        Legends &copy; Riot Games, Inc.
                    </li>
                </ul>
                <div className={styles.contact}>
                    <small>
                        <span>Contact:</span>&nbsp;
                        <a href="mailto:contact@controlward.com">
                            contact@controlward.com
                        </a>                        
                    </small>
                </div>
            </div>
        </footer>
    );
}

export default Footer;