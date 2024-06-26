
import { useState, useContext, useEffect } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { GlobalContext } from "../contexts/GlobalContext"
import searchResume from "../services/search/searchResume"
import startThread from "../services/chat/startThread"
import generateThreadName from "../services/chat/generateThreadName"

import { Tooltip } from "react-tooltip"
import ResumeSearchInput from "../components/ResumeSearchInput/ResumeSearchInput"
import IndeterminateCheckbox from "../components/IndeterminateCheckbox/IndeterminateCheckbox"
import SearchResultList from "../components/SearchResultList/SearchResultList"
import StackItem from "../components/StackItem/StackItem"
import DetailItemModal from "../components/DetailItemModal/DetailItemModal"
import LoadingModal from "../components/LoadingModal/LoadingModal"
import PageButtons from "../components/PageButtons/PageButtons"

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
    const resumeSearchPage = (searchParams.get('page') || 1) < 1 ? 1 : (searchParams.get('page') || 1);
    const [isStartChatSession, setIsStartChatSession] = useState(false);
    const [isNoResult, setIsNoResult] = useState(false);

    const [resultViewMode, setResultViewMode] = useState(ViewMode.LIST)

    const [searchResults, setSearchResults] = useState([]);
    const selectedCount = searchResults.filter(item => item.selected).length
    const isAllSelected = searchResults.every(item => item.selected)
    const isIndeterminate = searchResults.some(item => item.selected) && !isAllSelected

    // ====== Side Effects ======
    useEffect(() => {
        setSearchResults([]);
        searchResume(resumeSearchInput, resumeSearchLevel, resumeSearchPage)
            .then(res => {
                if (res === null) {
                    setIsNoResult(true);
                    setSearchResults([]);
                    return;
                }
                setIsNoResult(false);
                const updatedResults = res.map(item => ({ ...item, selected: globalContext.isItemSelected(item.id) }));
                setSearchResults(updatedResults);
            })
    }, [searchParams]);

    // ====== Event Handlers ======
    const resumeSearchKeyDownHandler = (e) => {
        if (e.key === 'Enter'
            && resumeSearchInput.trim() !== ''
            && (resumeSearchInput.trim() !== searchParams.get('query') || resumeSearchLevel !== searchParams.get('level'))) {
            navigate(`/search?query=${resumeSearchInput.trim()}&page=1&level=1`);
        }
    }
    const resumeSearchClickHandler = () => {
        if (resumeSearchInput.trim() !== ''
            && (resumeSearchInput.trim() !== searchParams.get('query') || resumeSearchLevel !== searchParams.get('level'))) {
            navigate(`/search?query=${resumeSearchInput.trim()}&page=1&level=1`);
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
        if (item.url !== "") {
            window.open(item.url, '_blank')
        }
        else {
            alert("No download link available")
        }
    }

    const stackItemDetailClickHandler = (item) => {
        globalContext.setDetailItem(item)
        globalContext.setShowDetailItemModal(true)
    }
    const stackItemRemoveClickHandler = (itemId) => {
        globalContext.popFromSelectedStack(itemId)
    }
    const startChatSessionHandler = () => {
        if (globalContext.selectedItemsStack.length === 0) {
            alert("Please select at least one item to start chat session");
            return;
        }

        setIsStartChatSession(true)
        const idsString = globalContext.selectedItemsStack.map(item => item.id).join(', ')
        const timeStr = generateThreadName();

        startThread(idsString, timeStr)
            .then(res => {
                if (res !== null) {
                    setIsStartChatSession(false);
                    globalContext.setSelectedItemsStack([]);
                    globalContext.setShowSelectedItemsStack(false);
                    globalContext.setSidebarThreads([
                        {
                            id: res.id,
                            name: timeStr,
                        },
                        ...globalContext.sidebarThreads
                    ]);
                    navigate(`/chat/${res.id}`);
                }
            });
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
    const detailItemModalDownloadHandler = () => {
        if (globalContext.detailItem.url !== "") {
            window.open(globalContext.detailItem.url, '_blank')
        }
        else {
            alert("No download link available")
        }
    }

    return (
        <main className="my-content-wrapper flex no-scrollbar">
            {/* ====== Search Result Window ====== */}
            <div className={`${globalContext.showSelectedItemsStack && 'xl:mr-72'} flex-1 transition-all duration-700 ease-in-out`}>
                {/* ====== Search Input ====== */}
                <div className="my-container-small pt-6">
                    <ResumeSearchInput
                        value={resumeSearchInput}
                        onChange={(e) => setResumeSearchInput(e.target.value)}
                        onPressEnter={resumeSearchKeyDownHandler}
                        onClickButton={resumeSearchClickHandler}
                    />
                </div>

                {/* ====== Search Slider ====== */}
                {/* <div className="my-container-small pt-3">
                    <ResumeSearchSlider
                        value={resumeSearchLevel}
                        onChange={(e) => setResumeSearchLevel(e.target.value)}
                    />
                </div> */}

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
                                    data-tooltip-id="add-list-tooltip"
                                    data-tooltip-content="Add selected items to List"
                                    data-tooltip-place="bottom"
                                    data-tooltip-delay-show={700}
                                >Add to List</button>
                                <Tooltip id="add-list-tooltip" className="z-50" />
                            </>
                        }
                    </div>

                    {/* ====== View Mode Buttons ====== */}
                    {/* <div className="flex items-center">
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
                    </div> */}
                </div>

                {/* ====== Search Results ====== */}
                <div className="my-container-medium min-h-[25rem] mt-4 pb-10">
                    {
                        isNoResult ? (
                            <p className="text-subtitle text-center">No result found</p>
                        )
                            :
                            (searchResults && searchResults.length === 0) ? (
                                <div className="mt-6 flex flex-col items-center space-y-4">
                                    <p className="text-subtitle">Loading search result ...</p>
                                    <div className="loader"></div>
                                </div>
                            ) : (
                                <SearchResultList
                                    searchResults={searchResults}
                                    viewMode={resultViewMode}
                                    onItemSelectClick={resultItemClickHandler}
                                    onItemDetailClick={resultItemDetailClickHandler}
                                    onItemDownloadClick={resultItemDownloadClickHandler}
                                />
                            )
                    }
                </div>

                {/* ====== Pagination ====== */}
                <div className="my-container-medium pb-12 flex justify-center">
                    {
                        searchResults && searchResults.length > 0 &&
                        <PageButtons
                            curr={resumeSearchPage}
                            query={searchParams.get('query')}
                            level={searchParams.get('level')}
                        />
                    }
                </div>
            </div>


            {/* ====== Selected Items Stack ====== */}
            <div className={`${globalContext.showSelectedItemsStack ? 'translate-x-0' : 'translate-x-full'} z-20 w-full max-w-72 h-[calc(100%-3rem)] fixed right-0 flex flex-col bg-background px-3 pt-3 pb-5 border-l-2 border-border transition-all duration-700 ease-in-out`}>
                <h1 className="text-lg font-semibold">Selected items ({globalContext.selectedItemsStack.length})</h1>

                <div className="flex-1 overflow-y-auto mt-3 mb-4">
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
                    onClick={startChatSessionHandler}
                    data-tooltip-id="start-chat-tooltip"
                    data-tooltip-content="Analyze profiles with AI Chatbot ✨"
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
                onDownloadClick={detailItemModalDownloadHandler}
            />

            {/* ====== Loading Modal ====== */}
            <LoadingModal showLoadingModal={isStartChatSession} />

        </main>
    )
}

export default SearchPage
