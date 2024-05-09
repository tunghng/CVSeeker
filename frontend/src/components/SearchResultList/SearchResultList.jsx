
import SearchResultItem from "../SearchResultItem/SearchResultItem";

const SearchResultList = ({ searchResults, setSearchResults, viewMode }) => {

    const handleItemClick = (id) => {
        const updatedResults = searchResults.map(item =>
            item.id === id ? { ...item, selected: !item.selected } : item
        );
        setSearchResults(updatedResults);
    }

    return (
        <ul>
            {searchResults.map(item => (
                <SearchResultItem
                    key={item.id}
                    item={item}
                    viewMode={viewMode}
                    handleItemClick={handleItemClick}
                />
            ))}
        </ul>
    );
}

export default SearchResultList
