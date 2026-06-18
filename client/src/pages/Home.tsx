import { usePreferences } from "../hooks/usePreferences"
// import styles from "./Home.module.css"

function Home() {
    const { preferences } = usePreferences()

    return <h1>Helo {preferences.theme}</h1>
}

export default Home
