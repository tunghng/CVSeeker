
import { useRef, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import FeatherIcon from 'feather-icons-react'

const HomePage = () => {
    // ====== State Management ======
    const [searchInput, setSearchInput] = useState('')
    const searchInputDOM = useRef(null)
    const navigate = useNavigate()

    // ====== Event Handlers ======
    const searchKeyDownHandler = (e) => {
        if (e.key === 'Enter') {
            navigate(`/search/${searchInput}`)
        }
    }

    return (
        <main>

            {/* ====== Search Input ====== */}
            <div className="my-container-small pt-6 relative flex items-center">
                <input
                    type="text"
                    className="flex-1 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-text font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out"
                    placeholder="Search..." 
                    value={searchInput}
                    onChange={(e) => setSearchInput(e.target.value)}
                    onKeyDown={searchKeyDownHandler}
                    required
                    autoComplete="true"
                    ref={searchInputDOM}
                />

                <Link to={`/search/${searchInput}`} className="absolute  right-10 sm:right-14 text-text peer-focus:text-primary transition-all duration-300 ease-in-out">
                    <FeatherIcon icon="search" className="w-6 h-6"/>
                </Link>
            </div>

        </main>
    )
}

export default HomePage
