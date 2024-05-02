
import { Link, useLocation } from "react-router-dom"
import './Sidebar.css'

const Sidebar = ({ showSidebar }) => {
    const location = useLocation()

    return (
        <div className={`${showSidebar ? 'translate-x-0' : '-translate-x-full'} w-60 h-full mt-12 fixed top-0 left-0 flex flex-col bg-background py-2 border-r-2 border-border transition-all duration-700 ease-in-out`}>

            <div className="navigation-list">
                <Link to='/' className='navigation-item'>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-search"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                    <span>New Search</span>
                </Link>

                <Link to='/saved' className={`navigation-item ${location.pathname === '/saved' && 'active'}`}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-bookmark"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path></svg>
                    <span>Saved CV</span>
                </Link>

                <Link to='/upload' className={`navigation-item ${location.pathname === '/upload' && 'active'}`}>
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-upload-cloud"><polyline points="16 16 12 12 8 16"></polyline><line x1="12" y1="12" x2="12" y2="21"></line><path d="M20.39 18.39A5 5 0 0 0 18 9h-1.26A8 8 0 1 0 3 16.3"></path><polyline points="16 16 12 12 8 16"></polyline></svg>
                    <span>Upload CV</span>
                </Link>
            </div>

            <div className="navigation-list">
                <div>
                    <h3 className='navigation-title'>Today</h3>
                    <Link to='/' className='navigation-item'>
                        <span>Find 10 CVs good at JavaScript</span>
                    </Link>
                    <Link to='/' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                </div>

                <div>
                    <h3 className='navigation-title'>24/04/2024</h3>
                    <Link to='/' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                    <Link to='/' className='navigation-item'>
                        <span>Find 8 people good at communicate</span>
                    </Link>
                </div>
            </div>


    </div>
    )
}

export default Sidebar
