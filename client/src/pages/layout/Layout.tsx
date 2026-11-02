import type { PropsWithChildren } from "react"
import styles from "./Layout.module.css"
import Header from "./Header"
import Footer from "./Footer"

type LayoutProps = PropsWithChildren<{
    hideHeader?: boolean
}>

function Layout({ children, hideHeader }: LayoutProps) {
    return <div className={styles.container}>
        <main className={styles.main + " " + (!hideHeader ? styles.withHeader : "")}>
            {children}
        </main>
        <Header hidden={hideHeader} />
        <Footer />
    </div>
}

export default Layout
