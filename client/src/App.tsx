import { BrowserRouter, Route, Routes } from "react-router"
import { PreferencesProvider } from "./hooks/usePreferences"
import Home from "./pages/Home"
import Profile from  "./pages/Profile" 
import About from "./pages/About"

export default function App() {
    return (
        <PreferencesProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/profile/:puuid" element={<Profile />} />
                    <Route path="/about" element={<About />} />
                </Routes>
            </BrowserRouter> 
        </PreferencesProvider>
    )
}





