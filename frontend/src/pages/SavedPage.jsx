
import { useRef, useState } from "react"

import IndeterminateCheckbox from "../components/IndeterminateCheckbox/IndeterminateCheckbox"

const SavedPage = () => {
    // ====== State Management ======
    const [viewMode, setViewMode] = useState('list')
    const [saveItems, setSaveItems] = useState([
        { id: 1, name: 'File 1', selected: false },
        { id: 2, name: 'File 2', selected: false },
        { id: 3, name: 'File 3', selected: false }
    ])
    const isAllSelected = saveItems.every(item => item.selected)
    const isIndeterminate = saveItems.some(item => item.selected) && !isAllSelected
    const selectedCount = saveItems.filter(item => item.selected).length

    // ====== Event Handlers ======
    const handleItemClick = (id) => {
        const updatedItems = saveItems.map(item =>
            item.id === id ? { ...item, selected: !item.selected } : item
        )
        setSaveItems(updatedItems)
    }

    const handleSelectAll = () => {
        const allSelected = saveItems.every(item => item.selected)
        const updatedItems = saveItems.map(item => ({ ...item, selected: !allSelected }))
        setSaveItems(updatedItems)
    }

    return (
        <main>
            <div className="my-container-medium py-6">
                <h1 className="text-2xl font-bold">Saved CV ({saveItems.length})</h1>
                
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
                            <button className="my-button my-button-outline">Unsave</button>
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

                {/* ====== Save Items ====== */}
                <div>
                    <ul>
                        {saveItems.map(item => (
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
        </main>
    )
}

export default SavedPage
