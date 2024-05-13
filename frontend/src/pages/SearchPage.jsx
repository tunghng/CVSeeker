
import { useState, useContext, useEffect } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { GlobalContext } from "../contexts/GlobalContext"

import search from "../services/search"

import { Tooltip } from "react-tooltip"
import FeatherIcon from 'feather-icons-react'
import ResumeSearchInput from "../components/ResumeSearchInput/ResumeSearchInput"
import ResumeSearchSlider from "../components/ResumeSearchSlider/ResumeSearchSlider"
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

    const navigate = useNavigate()

    const [searchParams, setSearchParams] = useSearchParams();
    const [resumeSearchInput, setResumeSearchInput] = useState(searchParams.get('query') || '');
    const [resumeSearchLevel, setResumeSearchLevel] = useState(searchParams.get('level') || 0.5);

    const [resultViewMode, setResultViewMode] = useState(ViewMode.LIST)

    const [searchResults, setSearchResults] = useState([]);
    const selectedCount = searchResults.filter(item => item.selected).length
    const isAllSelected = searchResults.every(item => item.selected)
    const isIndeterminate = searchResults.some(item => item.selected) && !isAllSelected

    // ====== Side Effects ======
    useEffect(() => {
        search(resumeSearchInput, resumeSearchLevel)
            .then(data => setSearchResults(data))
            .catch(error => console.error(error))
    }, [searchParams]);

    // ====== Event Handlers ======
    const resumeSearchKeyDownHandler = (e) => {
        if (e.key === 'Enter'
            && resumeSearchInput.trim() !== ''
            && (resumeSearchInput.trim() !== searchParams.get('query') || resumeSearchLevel !== searchParams.get('level'))) {
            navigate(`/search?query=${resumeSearchInput.trim()}&level=${resumeSearchLevel}`);
        }
    }
    const resumeSearchClickHandler = () => {
        if (resumeSearchInput.trim() !== ''
            && (resumeSearchInput.trim() !== searchParams.get('query') || resumeSearchLevel !== searchParams.get('level'))) {
            navigate(`/search?query=${resumeSearchInput.trim()}&level=${resumeSearchLevel}`);
        }
    }

    const resultSelectAllHandler = () => {
        const updatedItems = searchResults.map(item => ({ ...item, selected: !isAllSelected }))
        setSearchResults(updatedItems)
    }

    const addResultToStackHandler = () => {
        globalContext.pushToSelectedStack(searchResults.filter(item => item.selected))
        if (globalContext.showSelectedItemsStack === false) {
            globalContext.toggleSelectedItemsStack()
        }
    }

    const resultItemClickHandler = (id) => {
        const updatedResults = searchResults.map(item =>
            item.id === id ? { ...item, selected: !item.selected } : item
        );
        setSearchResults(updatedResults);
    }
    const resultItemDetailClickHandler = (item) => {
        globalContext.setDetailItem(item)
        globalContext.setShowDetailItemModal(true)
    }
    const resultItemDownloadClickHandler = (item) => {
        console.log(item)
    }

    const stackItemDetailClickHandler = (item) => {
        globalContext.setDetailItem(item)
        globalContext.setShowDetailItemModal(true)
    }
    const stackItemRemoveClickHandler = (itemId) => {
        globalContext.popFromSelectedStack(itemId)
    }

    const detailItemModalCloseHandler = () => {
        globalContext.setShowDetailItemModal(false)
    }
    const detailItemModalAddToListHandler = () => {
        globalContext.pushToSelectedStack([globalContext.detailItem])
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
                    <ResumeSearchInput
                        value={resumeSearchInput}
                        onChange={(e) => setResumeSearchInput(e.target.value)}
                        onPressEnter={resumeSearchKeyDownHandler}
                        onClickSearch={resumeSearchClickHandler}
                    />
                </div>

                {/* ====== Search Slider ====== */}
                <div className="my-container-small pt-3">
                    <ResumeSearchSlider
                        value={resumeSearchLevel}
                        onChange={(e) => setResumeSearchLevel(e.target.value)}
                    />
                </div>

                {/* ====== Actions Toolbar ====== */}
                <div className="my-container-medium flex justify-between mt-3 h-10">
                    {/* ====== Selecting Checkbox ====== */}
                    <div className="flex items-center gap-x-3 pl-3">
                        <IndeterminateCheckbox
                            checked={isAllSelected}
                            indeterminate={isIndeterminate}
                            onChange={resultSelectAllHandler}
                        />
                        <p>{selectedCount} selected</p>
                        {
                            (selectedCount > 0) &&
                            <>
                                <button
                                    className="my-button my-button-outline"
                                    onClick={addResultToStackHandler}
                                >Add to List</button>
                            </>
                        }
                    </div>

                    {/* ====== View Mode Buttons ====== */}
                    <div className="flex items-center">
                        <p className="mr-2">View as</p>
                        <button
                            className={`my-button my-button-outline-secondary px-3 rounded-l-full ${resultViewMode === ViewMode.LIST && 'bg-secondary-subtle hover:bg-secondary-subtle'}`}
                            onClick={() => setResultViewMode(ViewMode.LIST)}
                        >
                            <FeatherIcon icon="list" className="w-5 h-5" />
                        </button>
                        <button
                            className={`my-button my-button-outline-secondary px-3 rounded-r-full border-l-0 ${resultViewMode === ViewMode.GRID && 'bg-secondary-subtle hover:bg-secondary-subtle'}`}
                            onClick={() => setResultViewMode(ViewMode.GRID)}
                        >
                            <FeatherIcon icon="grid" className="w-5 h-5" />
                        </button>
                    </div>
                </div>

                {/* ====== Search Results ====== */}
                <div className="my-container-medium mt-4">
                    <SearchResultList
                        searchResults={searchResults}
                        viewMode={resultViewMode}
                        onItemSelectClick={resultItemClickHandler}
                        onItemDetailClick={resultItemDetailClickHandler}
                        onItemDownloadClick={resultItemDownloadClickHandler}
                    />
                </div>
            </div>


            {/* ====== Selected Items Stack ====== */}
            <div className={`${globalContext.showSelectedItemsStack ? 'translate-x-0' : 'translate-x-full'} w-full max-w-72 h-[calc(100%-3rem)] fixed  right-0 flex flex-col bg-background px-3 pt-3 pb-5 border-l-2 border-border transition-all duration-700 ease-in-out`}>
                <h1 className="text-lg font-semibold">Selected items ({globalContext.selectedItemsStack.length})</h1>

                <div className="flex-1">
                    {
                        globalContext.selectedItemsStack.map(item => (
                            <StackItem
                                key={item.id}
                                item={item}
                                onDetailClick={stackItemDetailClickHandler}
                                onRemoveClick={stackItemRemoveClickHandler}
                                showRemoveIcon={true}
                            />
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


            {/* ====== Detail Item Modal ====== */}
            <DetailItemModal
                showDetailItemModal={globalContext.showDetailItemModal}
                detailItem={globalContext.detailItem}
                onModalClose={detailItemModalCloseHandler}
                onAddToList={detailItemModalAddToListHandler}
            />

        </main>
    )
}

export default SearchPage
