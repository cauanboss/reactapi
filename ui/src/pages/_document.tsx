import { Html, Head, Main, NextScript } from "next/document";
import Link from "next/link";
let components;
export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <div className="table">
          <div className="table-row">
            <div className="table-fixed">
              <div className="table-auto">
                <Link href={"/home"}>Home</Link>
              </div>
              <div className="table-auto">
                <Link href={"/user"}>User</Link>
              </div>
            </div>
            <div className="table-cell">
              <Main />
            </div>
          </div>
        </div>
        <NextScript />
      </body>
    </Html>
  );
}
