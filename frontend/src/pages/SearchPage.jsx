
import { useState, useContext, useEffect } from "react"
import { useSearchParams } from "react-router-dom"
import { GlobalContext } from "../contexts/GlobalContext"

import search from "../services/search"

import { Tooltip } from "react-tooltip"
import FeatherIcon from 'feather-icons-react'
import SearchInput from "../components/SearchInput/SearchInput"
import SearchSlider from "../components/SearchSlider/SearchSlider"
import IndeterminateCheckbox from "../components/IndeterminateCheckbox/IndeterminateCheckbox"
import SearchResultList from "../components/SearchResultList/SearchResultList"
import StackItem from "../components/StackItem/StackItem"
import DetailItemModal from "../components/DetailItemModal/DetailItemModal"

const ViewMode = {
    GRID: 'grid',
    LIST: 'list'
};

const SearchPage = () => {
    // ====== State Management ======
    const globalContext = useContext(GlobalContext);

    const [searchParams, setSearchParams] = useSearchParams();
    const query = searchParams.get('query') || '';
    const level = searchParams.get('level') || 0.5;

    const [viewMode, setViewMode] = useState(ViewMode.LIST)

    const [searchResults, setSearchResults] = useState([]);
    const selectedCount = searchResults.filter(item => item.selected).length

    // ====== Fetching Data ======
    useEffect(() => {
        search(query, level)
            .then(data => setSearchResults(data))
            .catch(error => console.error(error))
    }, [query, level])

    // ====== Event Handlers ======
    const handleAddToList = () => {
        globalContext.pushToSelectedStack(searchResults.filter(item => item.selected))
        if (globalContext.showSelectedItemsStack === false) {
            globalContext.toggleSelectedItemsStack()
        }
    }

    return (
        <main className="h-full flex overflow-x-hidden">
            {/* ====== Search Result Window ====== */}
            <div className={`${globalContext.showSelectedItemsStack && 'md:mr-72'} flex-1 transition-all duration-700 ease-in-out`}>
                {/* ====== Search Input ====== */}
                <div className="my-container-small pt-6">
                    <SearchInput />
                </div>

                {/* ====== Search Slider ====== */}
                <div className="my-container-small pt-3">
                    <SearchSlider />
                </div>

                {/* ====== Actions Toolbar ====== */}
                <div className="my-container-medium flex justify-between mt-3 h-10">
                    {/* ====== Selecting Checkbox ====== */}
                    <div className="flex items-center gap-x-3 pl-3">
                        <IndeterminateCheckbox
                            searchResults={searchResults}
                            setSearchResults={setSearchResults}
                        />
                        <p>{selectedCount} selected</p>
                        {
                            (selectedCount > 0) &&
                            <>
                                <button
                                    className="my-button my-button-outline"
                                    onClick={handleAddToList}
                                >Add to List</button>
                            </>
                        }
                    </div>

                    {/* ====== View Mode Buttons ====== */}
                    <div className="flex items-center">
                        <p className="mr-2">View as</p>
                        <button
                            className={`my-button my-button-outline-secondary px-3 rounded-l-full ${viewMode === ViewMode.LIST && 'bg-secondary-subtle hover:bg-secondary-subtle'}`}
                            onClick={() => setViewMode(ViewMode.LIST)}
                        >
                            <FeatherIcon icon="list" className="w-5 h-5" />
                        </button>
                        <button
                            className={`my-button my-button-outline-secondary px-3 rounded-r-full border-l-0 ${viewMode === ViewMode.GRID && 'bg-secondary-subtle hover:bg-secondary-subtle'}`}
                            onClick={() => setViewMode(ViewMode.GRID)}
                        >
                            <FeatherIcon icon="grid" className="w-5 h-5" />
                        </button>
                    </div>
                </div>

                {/* ====== Search Results ====== */}
                <div className="my-container-medium mt-4">
                    <SearchResultList
                        searchResults={searchResults}
                        setSearchResults={setSearchResults}
                        viewMode={viewMode}
                    />
                </div>
            </div>


            {/* ====== Selected Items Stack ====== */}
            <div className={`${globalContext.showSelectedItemsStack ? 'translate-x-0' : 'translate-x-full'} w-full max-w-72 h-[calc(100%-3rem)] fixed  right-0 flex flex-col bg-background px-3 pt-3 pb-5 border-l-2 border-border transition-all duration-700 ease-in-out`}>
                <h1 className="text-lg font-semibold">Selected items ({globalContext.selectedItemsStack.length})</h1>

                <div className="flex-1">
                    {
                        globalContext.selectedItemsStack.map(item => (
                            <StackItem key={item.id} item={item} />
                        ))
                    }
                </div>

                <button
                    className="my-button my-button-primary py-2"
                    data-tooltip-id="start-chat-tooltip"
                    data-tooltip-content="Interact with AI chatbot in just a click!"
                    data-tooltip-place="top"
                    data-tooltip-delay-show={200}>
                    Start Chat Session
                </button>
                <Tooltip id="start-chat-tooltip" />

                {/* ====== Toggle Stack Button ====== */}
                <button
                    className="absolute top-1/2 -left-4 transform -translate-x-1/2 -translate-y-1/2"
                    onClick={globalContext.toggleSelectedItemsStack}
                >
                    <div className="flex h-12 w-6 flex-col items-center justify-center group"
                        data-tooltip-id="toggle-stack-tooltip"
                        data-tooltip-content={globalContext.showSelectedItemsStack ? 'Close Stackbar' : 'Open Stackbar'}
                        data-tooltip-place="left"
                        data-tooltip-delay-show={200}>
                        <div className={`${globalContext.showSelectedItemsStack ? 'rotate-0 group-hover:rotate-[-24deg]' : 'rotate-[24deg]'} selected-item-stack-button translate-y-[0.15rem]`}></div>
                        <div className={`${globalContext.showSelectedItemsStack ? 'rotate-0 group-hover:rotate-[24deg]' : 'rotate-[-24deg]'} selected-item-stack-button translate-y-[-0.15rem]`}></div>
                    </div>
                    <Tooltip id="toggle-stack-tooltip" />
                </button>

            </div>

            <DetailItemModal />
        </main>
    )
}

export default SearchPage
