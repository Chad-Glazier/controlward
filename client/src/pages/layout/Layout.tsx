import type { PropsWithChildren } from "react"
import styles from "./Layout.module.css"
import Header from "./Header"
import Footer from "./Footer"

type LayoutProps = PropsWithChildren<{
    hideHeader?: boolean
}>

function Layout({ children, hideHeader }: LayoutProps) {
    return <div className={styles.container}>
        <Header hidden={hideHeader} />
        <main className={styles.main + " " + (!hideHeader ? styles.withHeader : "")}>
            {children}
        </main>
        <Footer />
    </div>
}

export default Layout
