
import SearchResultItem from "../SearchResultItem/SearchResultItem";

const SearchResultList = ({ searchResults, viewMode, onItemSelectClick, onItemDetailClick, onItemDownloadClick }) => {

    return (
        <div className="w-full">
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
                        onSelectClick={onItemSelectClick}
                        onDetailClick={onItemDetailClick}
                        onDownloadClick={onItemDownloadClick}
                    />
                ))
            }
        </div>
    );
}

export default SearchResultList
