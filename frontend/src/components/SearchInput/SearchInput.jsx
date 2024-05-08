
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import FeatherIcon from 'feather-icons-react';

const SearchInput = () => {
    // ====== State Management ======
    const [searchInput, setSearchInput] = useState('')
    const navigate = useNavigate()

    // ====== Event Handlers ======
    const searchKeyDownHandler = (e) => {
        if (e.key === 'Enter' && searchInput.trim() !== ''){
            navigate(`/search/${searchInput}`)
        }
    }
    const searchClickHandler = () => {
        if (searchInput.trim() !== ''){
            navigate(`/search/${searchInput}`)
        }
    }

    return (
        <div className="relative flex items-center w-full">
            <input
                type="text"
                className="flex-1 pl-4 pr-11 py-2 peer bg-transparent rounded-full text-text font-medium text-lg outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out"
                placeholder="Search..."
                value={searchInput}
                onChange={(e) => setSearchInput(e.target.value)}
                onKeyDown={searchKeyDownHandler}
            />

            <button
                onClick={searchClickHandler}
                className="absolute right-2 p-2 rounded-full text-text peer-focus:text-primary transition-all duration-300 ease-in-out"
            >
                <FeatherIcon icon="search" className="w-6 h-6" />
            </button>
        </div>
    );
};

export default SearchInput;
