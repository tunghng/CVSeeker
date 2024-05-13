
import { Link, useLocation } from "react-router-dom"
import { useContext, useState } from "react"
import { GlobalContext } from "../../contexts/GlobalContext"

import FeatherIcon from 'feather-icons-react'
import SidebarThreadItem from "../SidebarThreadItem/SidebarThreadItem"

import './Sidebar.css'

const Sidebar = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    const location = useLocation()

    const [threadList, setThreadList] = useState([
        { id: '123', name: 'Find 10 CVs good at JavaScript' },
        { id: '456', name: 'Find 8 people good at communicate' },
        { id: '234', name: 'Find 5 people good at communicate' },
        { id: '102', name: 'Find 4 people good at communicate' },
    ])

    return (
        <div className={`${globalContext.showSidebar ? 'translate-x-0' : '-translate-x-full'} w-64 h-full mt-12 fixed top-0 left-0 flex flex-col z-10 bg-background py-2 border-r-2 border-border transition-all duration-700 ease-in-out`}>

            {/* ====== Page Navigation List ====== */}
            <div className="navigation-list">
                <Link to='/' className='navigation-item'>
                    <FeatherIcon icon="search" className="w-6 h-6" strokeWidth={1.8} />
                    <span>New Search</span>
                </Link>

                <Link to='/saved' className={`navigation-item ${location.pathname === '/saved' && 'active'}`}>
                    <FeatherIcon icon="bookmark" className="w-6 h-6" strokeWidth={1.8} />
                    <span>Saved CV</span>
                </Link>

                <Link to='/upload' className={`navigation-item ${location.pathname === '/upload' && 'active'}`}>
                    <FeatherIcon icon="upload" className="w-6 h-6" strokeWidth={1.8} />
                    <span>Upload CV</span>
                </Link>
            </div>


            {/* ====== Thread List ====== */}
            <div className="thread-list">
                <h3 className='thread-title mt-6'>Threads</h3>
                <div>
                    {threadList.map((item, index) => (
                        <SidebarThreadItem
                            key={index}
                            item={item}
                            isActive={location.pathname === `/chat/${item.id}`}
                        />
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Sidebar
