
import { useRef, useState } from "react"
import { Link, useNavigate } from "react-router-dom"

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
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-search"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                </Link>
            </div>

        </main>
    )
}

export default HomePage
