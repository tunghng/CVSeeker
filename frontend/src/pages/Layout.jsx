import { Link, Outlet } from "react-router-dom"

const Layout = () => {

    return (
        <>
            {/* ====== Header Navigation ====== */}
            <nav>
                <Link to='/'>Home</Link>
                <Link to='/upload'>Upload CV</Link>
            </nav>

            <Outlet />
        </>
    )
}

export default Layout