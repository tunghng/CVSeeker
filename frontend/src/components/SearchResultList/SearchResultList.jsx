
import SearchResultItem from "../SearchResultItem/SearchResultItem";

const SearchResultList = ({ searchResults, setSearchResults, viewMode }) => {

    const handleItemClick = (id) => {
        const updatedResults = searchResults.map(item =>
            item.id === id ? { ...item, selected: !item.selected } : item
        );
        setSearchResults(updatedResults);
    }

    return (
        <div className="w-full mt-4">
            <div className="h-10 flex border-b border-border">
                <div className="w-10"></div>
                <div className="flex-1 flex items-center text-left font-bold">Name</div>
                <div className="w-32 flex items-center justify-center font-bold">Imported Date</div>
                <div className="w-24 sm:w-56"></div>
            </div>

            {searchResults.length === 0 &&
                <div>
                    <p className="text-center py-4">No results found</p>
                </div>
            }

            {searchResults.length > 0 &&
                searchResults.map(item => (
                    <SearchResultItem
                        key={item.id}
                        item={item}
                        viewMode={viewMode}
                        handleItemClick={handleItemClick}
                    />
                ))
            }
        </div>
    );
}

export default SearchResultList
