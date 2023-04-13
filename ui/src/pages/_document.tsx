import { Html, Head, Main, NextScript } from 'next/document'
import Link from 'next/link'
let components
export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <div style={{display: 'inline-block', width: '100%', height: '100%'}}>
          <div style={{width: '200px', height: '100%', float: 'left'}}>
            <Link href={'/home'}>Home</Link>
            <br />
            <Link href={'/user'}>User</Link>
          </div>
          <div >
            <Main />
          </div>
        </div>
        <NextScript />
      </body>
    </Html>
  )
}
