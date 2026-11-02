import Layout from "./layout/Layout"
import styles from "./Profile.module.css"
import { useParams } from "react-router"

function Profile() {
    const { puuid } = useParams()

    return <Layout>
        <div className={styles.placeholder}>

        </div>
        {puuid}
    </Layout>
}

export default Profile
