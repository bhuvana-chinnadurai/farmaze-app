"use client"
import { useEffect, useState } from "react";
import Image from "next/image";
import { usePathname, useRouter } from "next/navigation";
import Link from "next/link";

const withHeader = ComponentToRender => (props) => {
  const [isLoggedIn, setLoggedInStatus] = useState(false);
  const pathName = usePathname();
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem("app-token");
    if (!token && pathName !== "/login") router.push("/login");
    else if (!!token) {
      setLoggedInStatus(true);
      if (pathName === "/login") router.push("/");
    }
  }, []);

  return (
    <>
      <div className='header'>
        <div className='header-left'>
          <Link href="/"><Image src="/logo.png" width={140} height={44} alt="Farmaze Logo" /></Link>
        </div>
        <div className='header-right'>
          {/* <div><Image src="/contact.png" width={32} height={32} alt="Contact Support" /></div> */}
          {/* <AntdButton height={48} type="primary">Login</AntdButton> */}
        </div>
      </div>
      <ComponentToRender {...props} isLoggedIn={isLoggedIn} />
    </>
  )
}

export default withHeader;
