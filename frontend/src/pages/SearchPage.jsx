
import { useRef, useState } from "react"
import { Link, useNavigate, useParams } from "react-router-dom"

import logocvseeker from '../assets/images/logo.png'
import IndeterminateCheckbox from "../components/IndeterminateCheckbox/IndeterminateCheckbox"

const SearchPage = () => {
    // ====== State Management ======
    const { id } = useParams()
    const [viewMode, setViewMode] = useState('grid')
    const [searchItems, setSearchItems] = useState([
        { id: 1, name: 'File 1', selected: false },
        { id: 2, name: 'File 2', selected: false },
        { id: 3, name: 'File 3', selected: false }
    ])
    const isAllSelected = searchItems.every(item => item.selected)
    const isIndeterminate = searchItems.some(item => item.selected) && !isAllSelected
    const selectedCount = searchItems.filter(item => item.selected).length
    const [showChatWindow, setShowChatWindow] = useState(false)

    // ====== Event Handlers ======
    const handleItemClick = (id) => {
        const updatedItems = searchItems.map(item =>
            item.id === id ? { ...item, selected: !item.selected } : item
        )
        setSearchItems(updatedItems)
    }

    const handleSelectAll = () => {
        const allSelected = searchItems.every(item => item.selected)
        const updatedItems = searchItems.map(item => ({ ...item, selected: !allSelected }))
        setSearchItems(updatedItems)
    }


    return (
        <main className="h-full flex overflow-x-hidden">
            {/* ====== Search Result Window ====== */}
            <div className="flex-1">
                {/* ====== Search Input ====== */}
                <div className="my-container-small pt-6 relative flex items-center">
                    <input
                        type="text"
                        className="flex-1 pl-4 pr-11 py-2 bg-disable-light rounded-full text-subtitle font-medium text-lg outline-none border-2 border-disable-light"
                        placeholder="Search..." 
                        value={id}
                        disabled
                    />
                    <Link to={`/search/${id}`} className="absolute right-10 sm:right-14 text-subtitle pointer-events-none">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-search"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                    </Link>
                </div>

                <div className="my-container-medium">
                    {/* ====== Actions Toolbar ====== */}
                    <div className="flex justify-between mt-3 h-10">
                        {/* ====== Selecting Checkbox ====== */}
                        <div className="flex items-center gap-x-3">
                            <IndeterminateCheckbox
                                checked={isAllSelected}
                                indeterminate={isIndeterminate}
                                onChange={handleSelectAll}
                            />
                            <p>{selectedCount} selected</p>
                            {
                                (selectedCount > 0) &&
                                <button className="my-button my-button-outline">Save</button>
                            }
                        </div>

                        {/* ====== View Mode Buttons ====== */}
                        <div className="flex items-center">
                            <p className="mr-2">View as</p>
                            <button 
                                className={`my-button my-button-outline-secondary px-3 rounded-l-full ${viewMode === 'list' && 'bg-primary-subtle hover:bg-primary-subtle'}`}
                                onClick={() => setViewMode('list')}
                                >
                                <svg className="feather feather-list" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg>
                            </button>
                            <button
                                className={`my-button my-button-outline-secondary px-3 rounded-r-full border-l-0 ${viewMode === 'grid' && 'bg-primary-subtle hover:bg-primary-subtle'}`}
                                onClick={() => setViewMode('grid')}
                                >
                                <svg className="feather feather-grid" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
                            </button>
                        </div>
                    </div>

                    {/* ====== Search Results ====== */}
                    <div>
                        <ul>
                            {searchItems.map(item => (
                                <li key={item.id} onClick={() => handleItemClick(item.id)}>
                                    <input
                                        type="checkbox"
                                        checked={item.selected}
                                        readOnly
                                    />
                                    <label>{item.name}</label>
                                </li>
                            ))}
                        </ul>
                    </div>
                </div>
            </div>

            {/* ====== Chat Window ====== */}
            <div className={`${showChatWindow ? 'w-full' : 'w-0'} relative max-w-[30rem] border-l-2 border-border flex flex-col`}>

                {/* ====== Start Chat Button ====== */}
                <div className="start-chat-button-container">
                    <div className="start-chat-button-wrapper drop-shadow-lg">
                        <button className="start-chat-button" onClick={() => setShowChatWindow(!showChatWindow)}>
                            <img src={logocvseeker} alt="logo" className="h-10 w-10 p-1" />
                            <p className="text-primary font-semibold">Start chat session</p>
                        </button>
                    </div>
                </div>

                <div className="flex-1">Chat messages</div>
                
                <div className="flex px-4 py-6 gap-x-2">
                    <input 
                        type="text" 
                        placeholder="Type a message..." 
                        className="flex-1 px-2 py-2 rounded-lg text-text text-base outline-none border-2 border-border focus:border-primary transition-all duration-300 ease-in-out"
                    />
                    <button className="my-button my-button-subtle">
                        <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="feather feather-send"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                    </button>
                </div>
            </div>
        </main>
    )
}

export default SearchPage