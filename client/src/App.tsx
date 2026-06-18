import { BrowserRouter, Route, Routes } from "react-router"
import Home from "./pages/Home"
import { PreferencesProvider } from "./hooks/usePreferences"

export default function App() {
    return (
        <PreferencesProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home />} />
                </Routes>
            </BrowserRouter> 
        </PreferencesProvider>
    )
}





