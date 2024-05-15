
import { useContext } from "react"
import { GlobalContext } from "../contexts/GlobalContext"
import { Outlet } from "react-router-dom"

import Header from '../components/Header/Header'
import Sidebar from '../components/Sidebar/Sidebar'

const Layout = () => {
    const globalContext = useContext(GlobalContext);

    return (
        <div className="flex flex-col h-screen">
            {/* ====== Header ====== */}
            <Header />

            <div className="flex-1 flex">
                {/* ====== SideBar Navigation ====== */}
                <Sidebar />

                {/* ====== Child route rendering ====== */}
                <div className={`${globalContext.showSidebar && 'md:ml-64'} flex-1 mt-12 bg-background transition-all duration-700 ease-in-out`}>
                    <Outlet />
                </div>
            </div>
        </div>
    )
}

export default Layout
