import * as React from "react"
import { Link } from "gatsby"

import Layout from "../components/layout"
import Seo from "../components/seo"

import helloGIF from "../images/gif/cute.gif"

const IndexPage = () => (
  <Layout>
    <Seo title="Home" />
    <h3>hello &#128075;,</h3>
    <p>
      my name is Vincent and this is my personal website. There are many like it
      but this one is mine.
    </p>
    <p>
      I believe in the power and merit of open source software and a free and
      open internet.
    </p>
    <p>
      Feel free to write me <a href="mailto:hi@vinckr.com">an email</a> about
      anything, I would love to hear from you!
    </p>
    <p>
      Have a look around, I collect interesting <Link to="/links">links</Link>{" "}
      and hack away at a <Link to="/blog">blog</Link>.
    </p>
    <p>🗿🗿🗿🗿🗿 </p>
    <img src={helloGIF} alt="Greetings" />
  </Layout>
)

export default IndexPage
