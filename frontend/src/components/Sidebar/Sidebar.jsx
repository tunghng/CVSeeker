
import { Link, useLocation } from "react-router-dom"
import { useContext } from "react"
import { GlobalContext } from "../../contexts/GlobalContext"

import FeatherIcon from 'feather-icons-react'

import './Sidebar.css'

const Sidebar = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    const location = useLocation()

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

            {/* ====== Chat Navigation List ====== */}
            <div className="navigation-list">
                <div>
                    <h3 className='navigation-title'>Today</h3>
                    <Link to='/chat/123' className='navigation-item'>
                        <span>Find 10 CVs good at JavaScript</span>
                    </Link>
                    <Link to='/chat/456' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                </div>

                <div>
                    <h3 className='navigation-title'>24/04/2024</h3>
                    <Link to='/chat/hehe' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                    <Link to='/chat/abc' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                </div>
            </div>


        </div>
    )
}

export default Sidebar
