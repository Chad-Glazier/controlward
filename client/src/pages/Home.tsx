import Layout from "./layout/Layout"
import styles from "./Home.module.css"
import PlayerSearch from "../components/PlayerSearch"

function Home() {
    return <Layout hideHeader>
        <div className={styles.container}>
            <h1 className={styles.header}>Control<br />Ward</h1>
            <div className={styles.searchSection}>
                <p className={styles.subtitle}>Simple and performant League of Legends stat tracking</p>
                <PlayerSearch 
                    className={styles.searchBar}
                    onSearch={async (a, b) => {
                        await new Promise(resolve => setTimeout(resolve, 1000))
                        console.log(a, b)
                    }}
                />
                <div className={styles.recentSearches}>
                    <h2>Recently viewed</h2>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                    <div style={{
                        width: "100%",
                        minHeight: "100px",
                        backgroundColor: "#444",
                    }}></div>
                </div>
            </div>
        </div>
    </Layout>
}

export default Home
 