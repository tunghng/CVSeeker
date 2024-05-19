
import { useContext } from "react"
import { GlobalContext } from "../../contexts/GlobalContext"
import FeatherIcon from 'feather-icons-react'

import logocvseeker from '../../assets/images/logo.png'

const Header = () => {
    const globalContext = useContext(GlobalContext);

    return (
        <div className="h-12 px-3 bg-background flex justify-start items-center border-b border-border">

            {/* ====== Toggle Sidebar Button ====== */}
            <button
                className="my-button my-button-icon my-button-outline"
                onClick={globalContext.toggleSidebar}
            >
                {
                    globalContext.showSidebar ?
                        <FeatherIcon icon="chevrons-left" className="w-6 h-6" strokeWidth={1.9}/>
                        :
                        <FeatherIcon icon="menu" className="w-6 h-6" strokeWidth={1.9}/>
                }
            </button>

            {/* ====== Logo Image ====== */}
            <div className='absolute left-1/2 transform -translate-x-1/2 flex items-center select-none'>
                <img
                    src={logocvseeker}
                    alt="logo"
                    className="h-9"
                />
                <h1 className='text-xl text-primary font-bold'>CV Seeker</h1>
            </div>
        </div>
    )
}

export default Header
