import { BrowserRouter, Route, Routes } from "react-router"
import { PreferencesProvider } from "./hooks/usePreferences"
import Home from "./pages/Home"
import Profile from  "./pages/Profile" 

export default function App() {
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





