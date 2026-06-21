import { usePreferences } from "../hooks/usePreferences"
// import styles from "./Home.module.css"

function Home() {
    const { preferences, updatePreference } = usePreferences()

    return <>
        <p>theme: {preferences.theme}</p>
        <button
            onClick={() => {
                switch (preferences.theme) {
                case "dark":
                    updatePreference("theme", "light")
                    break
                case "light":
                    updatePreference("theme", "dark")
                    break
                }
            }}
        >
            toggle theme
        </button>
    </>
}

export default Home
 