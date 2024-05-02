
import { useState } from "react"
import { Outlet } from "react-router-dom"

import Header from '../components/Header/Header'
import Sidebar from '../components/Sidebar/Sidebar'

const Layout = () => {
    const [showSidebar, setShowSidebar] = useState(true)

    return (
        <div className="flex flex-col h-screen">
            {/* ====== Header ====== */}
            <Header showSidebar={showSidebar} setShowSidebar={setShowSidebar} />

            <div className="flex-1 flex">
                {/* ====== SideBar Navigation ====== */}
                <Sidebar showSidebar={showSidebar} />

                {/* ====== Child route rendering ====== */}
                <div className={`${showSidebar && 'md:ml-60'} flex-1 h-full bg-background overflow-y-scroll transition-all duration-700 ease-in-out`}>
                    <Outlet />
                </div>
            </div>
        </div>
    )
}

export default Layout
