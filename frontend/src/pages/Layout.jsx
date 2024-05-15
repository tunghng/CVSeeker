
import { useContext, useEffect, useState } from "react"
import { GlobalContext } from "../contexts/GlobalContext"
import { Outlet } from "react-router-dom"
import getAllThreads from "../services/chat/getAllThreads"

import Header from '../components/Header/Header'
import Sidebar from '../components/Sidebar/Sidebar'

const Layout = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);

    const [threads, setThreads] = useState([]);

    // ====== Fetching Threads ======
    useEffect(() => {
        getAllThreads()
            .then(res => {
                setThreads(res);
            })
    }, [])

    return (
        <div className="flex flex-col h-screen">
            {/* ====== Header ====== */}
            <Header />

            <div className="flex-1 flex">
                {/* ====== SideBar Navigation ====== */}
                <Sidebar threads={threads} />

                {/* ====== Child route rendering ====== */}
                <div className={`${globalContext.showSidebar && 'md:ml-64'} flex-1 mt-12 bg-background transition-all duration-700 ease-in-out`}>
                    <Outlet />
                </div>
            </div>
        </div>
    )
}

export default Layout
