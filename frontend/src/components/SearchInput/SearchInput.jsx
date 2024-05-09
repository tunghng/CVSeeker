
import { useState, useContext, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom"
import { GlobalContext } from "../../contexts/GlobalContext"

import FeatherIcon from 'feather-icons-react';

const SearchInput = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);
    const [searchParams, setSearchParams] = useSearchParams();
    const [searchInput, setSearchInput] = useState('');
    const [searchLevel, setSearchLevel] = useState(0.5);

    const navigate = useNavigate()

    useEffect(() => {
        const query = searchParams.get('query') || '';
        const level = searchParams.get('level') || 0.5;
        setSearchInput(query);
        setSearchLevel(level);
    }, [searchParams]);

    // ====== Event Handlers ======
    const searchKeyDownHandler = (e) => {
        if (e.key === 'Enter' 
            && searchInput.trim() !== '' 
            && (searchInput.trim() !== searchParams.get('query') || globalContext.sliderValue !== searchLevel)) {
            navigate(`/search?query=${searchInput.trim()}&level=${globalContext.sliderValue}`);
        }
    }
    const searchClickHandler = () => {
        if (searchInput.trim() !== '' 
            && (searchInput.trim() !== searchParams.get('query') || globalContext.sliderValue !== searchLevel)) {
            navigate(`/search?query=${searchInput.trim()}&level=${globalContext.sliderValue}`);
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
