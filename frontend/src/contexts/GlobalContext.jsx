
import { useState, createContext } from 'react';

const GlobalContext = createContext();

function GlobalProvider({ children }) {

    // ====== Left Sidebar state
    const [showSidebar, setShowSidebar] = useState(true);
    
    const toggleSidebar = () => {
        setShowSidebar(!showSidebar);
    }



    const value = {
        showSidebar,
        toggleSidebar
    }

    return (
        <GlobalContext.Provider value={value}>
            {children}
        </GlobalContext.Provider>
    )
}

export { GlobalContext, GlobalProvider }