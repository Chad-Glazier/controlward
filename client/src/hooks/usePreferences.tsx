import { createContext, useContext, useEffect, useState } from "react"

type Preferences = {
    theme: "light" | "dark"
    region: "americas"
}

const PREFERENCES_DEFAULT: Preferences = {
    theme: "dark",
    region: "americas"
}

type PreferencesContextType = {
    preferences: Preferences
    updatePreference: 
        <K extends keyof Preferences>(k: K, v: Preferences[K]) => void
}

const PreferencesContext = createContext<PreferencesContextType | null>(null)

export function PreferencesProvider({ children }: React.PropsWithChildren) {
    const [preferences, setPreferences] = useState<Preferences>(() => {
        const stored = localStorage.getItem("preferences")
        if (stored == null) {
            return PREFERENCES_DEFAULT
        }

        try {
            return JSON.parse(stored) as Preferences
        } catch {
            return PREFERENCES_DEFAULT
        }
    })

    useEffect(() => {
        localStorage.setItem("preferences", JSON.stringify(preferences))
    }, [preferences])

    function updatePreference<K extends keyof Preferences>(
        key: K, 
        value: Preferences[K]
    ) {
        setPreferences(prev => {
            return {
                ...prev,
                [key]: value,
            }
        })
    }

    return <PreferencesContext value={{ preferences, updatePreference }}>
        {children}
    </PreferencesContext>
}

export function usePreferences() {
    const ctx = useContext(PreferencesContext)
    
    if (!ctx) {
        throw new Error(
            "usePreferences must be used within a child of PreferencesContext"
        )
    }

    return ctx
}
