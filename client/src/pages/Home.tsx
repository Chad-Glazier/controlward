import Layout from "./layout/Layout"
import styles from "./Home.module.css"
import PlayerSearch from "../components/PlayerSearch"
import Logo from "../components/Logo"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faAngleDown  } from "@fortawesome/free-solid-svg-icons"
import { useState } from "react"
import { useNavigate } from "react-router"

function Home() {

    const navigate = useNavigate()

    let recents = [1,2,3,4,5,6,7,8,9,10]

    const [recentsToShow, setRecentsToShow] = useState(Math.min(4, recents.length))

    return <Layout hideHeader>
        <div className={styles.container}>
            <Logo height="24vmin" />
            <div className={styles.searchSection}>
                <p className={styles.subtitle}>Simple and performant League of Legends stat tracking</p>
                <PlayerSearch 
                    className={styles.searchBar}
                    onSearch={async (a) => {
                        navigate(`/profile/${a}`)
                    }}
                />
                <div className={styles.recentSearches}>
                    <h2>Recently viewed</h2>
                    {recents.slice(0, recentsToShow).map(() => 
                        <div style={{
                            width: "100%",
                            minHeight: "80px",
                            backgroundColor: "#444",
                        }}></div>                    
                    )}
                    {recentsToShow < recents.length && <button 
                        className={styles.downButton}
                        onClick={() => {
                            setRecentsToShow(prev => Math.min(prev + 10, recents.length))
                        }}    
                    >
                        <FontAwesomeIcon icon={faAngleDown} />
                    </button>}
                </div>
            </div>
        </div>
    </Layout>
}

export default Home
 