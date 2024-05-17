
import SearchResultItem from "../SearchResultItem/SearchResultItem";

const SearchResultList = ({ searchResults, viewMode, onItemSelectClick, onItemDetailClick, onItemDownloadClick }) => {

    return (
        <div className="w-full">
            <div className="h-10 flex border-b border-border sticky top-0 bg-slate-300 z-10">
                <div className="w-12"></div>
                <div className="w-44 flex items-center text-left font-bold">Full Name</div>
                <div className="w-32 flex items-center justify-center font-bold">Education</div>
                <div className="flex-1 flex items-center justify-center font-bold">Majors</div>
                <div className="w-20 sm:w-48"></div>
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
