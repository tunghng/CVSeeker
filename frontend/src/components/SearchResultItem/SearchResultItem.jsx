
const SearchResultItem = ({ item, viewMode, handleItemClick }) => {

    return (
        <li onClick={() => handleItemClick(item.id)}>
            <input
                type="checkbox"
                checked={item.selected}
                readOnly
            />
            <label>{item.name}</label>
        </li>
    );
}

export default SearchResultItem
