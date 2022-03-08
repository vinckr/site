import * as React from "react"
import { Link } from "gatsby"

import Layout from "../components/layout"
import Seo from "../components/seo"

import helloGIF from '../images/gif/cute.gif'

const IndexPage = () => (
  <Layout>
    <Seo title="Home" />
    <h3>Welcome! &#128075;</h3>
    <p>
     My name is Vincent and this is my personal website. There are many like it but this one is mine. 
  </p>
  <p>
    I believe in the power and merit of open source software and a free and open internet. 
  </p>
  <p>
    I work at <Link to="https://ory.sh/">Ory</Link> building a resilient developer community.
  </p>
  <p>
    Feel free to <a href="mailto:hi@vinckr.com">contact me</a>, if you have a request, question, or need help with your open source project. </p>
  <p>
    Have a look around, I have an old-timey <a href="./links.html">link collection</a>.
  </p>
  <img src={helloGIF} alt="Greetings" />
  </Layout>
)

export default IndexPage
