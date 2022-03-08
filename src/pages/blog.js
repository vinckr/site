import * as React from "react"
import { Link, graphql, StaticQuery } from "gatsby"

import Layout from "../components/layout"
import Seo from "../components/seo"

const BlogIndex = ({ location }) => {
  const data = useStaticQuery(graphql`
query {
    allMarkdownRemark(filter: {fileAbsolutePath: {regex: "/blog/"}}) {
      edges {
        node {
          id
          frontmatter {
            title
            slug
            date
          }
        }
      }
    }
  }
`
)

const posts = data.markdownRemark.edges.node

return (
    <Layout location={location}>
      <Seo title="All posts" />
      <ol style={{ listStyle: `none` }}>
        {posts.map(post => {
          const title = post.frontmatter.title || post.frontmatter.slug

          return (
            <li key={post.frontmatter.slug}>
              <article
                className="post-list-item"
                itemScope
                itemType="http://schema.org/Article"
              >
                <header>
                  <h2>
                    <Link to={post.frontmatter.slug} itemProp="url">
                      <span itemProp="headline">{title}</span>
                    </Link>
                  </h2>
                  <small>{post.frontmatter.date}</small>
                </header>
                <section>
                  <p
                    dangerouslySetInnerHTML={{
                      __html: post.frontmatter.description || post.excerpt,
                    }}
                    itemProp="description"
                  />
                </section>
              </article>
            </li>
          )
        })}
      </ol>
    </Layout>
  )
}

export default BlogIndex
