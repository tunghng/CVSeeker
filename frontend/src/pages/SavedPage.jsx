
import { useRef, useState } from "react"
import FeatherIcon from 'feather-icons-react'

import IndeterminateCheckbox from "../components/IndeterminateCheckbox/IndeterminateCheckbox"

const ViewMode = {
    GRID: 'grid',
    LIST: 'list'
};

const SavedPage = () => {
    // ====== State Management ======
    const [viewMode, setViewMode] = useState(ViewMode.LIST)
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
                            className={`my-button my-button-outline-secondary px-3 rounded-l-full ${viewMode === ViewMode.LIST && 'bg-primary-subtle hover:bg-primary-subtle'}`}
                            onClick={() => setViewMode(ViewMode.LIST)}
                            >
                            <FeatherIcon icon="list" className="w-5 h-5"/>
                        </button>
                        <button
                            className={`my-button my-button-outline-secondary px-3 rounded-r-full border-l-0 ${viewMode === ViewMode.GRID && 'bg-primary-subtle hover:bg-primary-subtle'}`}
                            onClick={() => setViewMode(ViewMode.GRID)}
                            >
                            <FeatherIcon icon="grid" className="w-5 h-5"/>
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
