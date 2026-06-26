import styles from "./About.module.css"
import Layout from "./layout/Layout"

function About() {
    return <Layout>
        <article className={styles.container}>
            <section>
                <h1>What is Control Ward?</h1>
                <p>
                    Control Ward is a minimalist stat tracking app for League of
                    Legends. 
                </p>                
            </section>
            <section>
                <h1>Why is Control Ward?</h1>
                <p>
                    You might be wondering, &ldquo;why make another OP.GG
                    clone?&rdquo; The answer to that question is twofold: 
                    <ol>
                        <li>
                            First, I feel that the quality of existing League stat 
                            trackers is sub-par. I'm not going to elaborate on this
                            because I don't want it to sound like I'm throwing shade.
                        </li>
                        <li>
                            The second, much more important reason is that I'm 
                            dumb enough to think I can do better.
                        </li>
                    </ol>
                </p>
            </section>
            <section>
                <h1>How is Control Ward?</h1>
                <p>
                    Control Ward has a typical server and client, both of which are in
                    a monorepo you can find&nbsp;
                    <a href="http://github.com/Chad-Glazier/controlward" target="_blank">here</a>.
                    The server is written in Go and acts as a layer between the client 
                    and the Riot API. In an effort to minimize the number of calls to
                    Riot, certain immutable data is cached by the server. The most important
                    cache is the match cache&mdash;since finished matches will never change,
                    the server can freely store them without the risk of holding outdated data. 
                    Additionally, unlike some other stat trackers, all large packets sent to/from 
                    the Control Ward server are compressed to improve network performance.
                </p>
                <p>
                    The client is the thing you're seeing right now. For the most part, it's 
                    a simple static site made in React. The cool part is that all data analytics 
                    are computed by a WebAssembly program running on the browser. This lets the 
                    website load analytics quickly and without eating a gigabyte of RAM or slowing 
                    down the server.
                </p>
            </section>
            <section>
                <h1>Who is Control Ward?</h1>
                <p>
                    Control Ward is made by me, <a target="_blank" href="https://chadglazier.com/">Chad Glazier</a>,
                    but it couldn't exist without the work that Riot has put into making
                    public APIs and giving free access to their static assets. They're 
                    very cool for that.
                </p>
            </section>
        </article>
    </Layout>
}

export default About
