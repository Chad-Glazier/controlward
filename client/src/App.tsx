import { BrowserRouter, Route, Routes } from "react-router"
import { PreferencesProvider } from "./hooks/usePreferences"
import { useEffect } from "react"
import Home from "./pages/Home"
import Profile from  "./pages/Profile" 

export default function App() {
    useEffect(() => {
        alert("ControlWard is currently in early development and has not yet been approved by Riot.")
    })

    return (
        <PreferencesProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/profile/:puuid" element={<Profile />} />
                </Routes>
            </BrowserRouter> 
        </PreferencesProvider>
    )
}
