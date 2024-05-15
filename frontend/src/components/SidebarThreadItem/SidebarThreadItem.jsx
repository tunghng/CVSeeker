
import { Link } from "react-router-dom"

import FeatherIcon from "feather-icons-react"
import { Tooltip } from "react-tooltip";

const SidebarThreadItem = ({ item, isActive }) => {
    return (
        <Link to={`/chat/${item.id}`}
            className={`thread-item group ${isActive && 'active'}`}
        >
            <span>{item.name === '' ? 'No Name' : item.name}</span>

            <button
                className="ml-3 rounded-md hidden group-hover:block hover:opacity-80 transition-all duration-300 ease-in-out"
                data-tooltip-id="item-tooltip"
                data-tooltip-content="Show selected items of this thread"
                data-tooltip-place="top"
                data-tooltip-delay-show={400}
            >
                <FeatherIcon icon="inbox" className="w-8 h-8 p-1" strokeWidth={1.8} />
            </button>
            <Tooltip id="item-tooltip" className="hidden group-hover:block" />
        </Link>
    )
}

export default SidebarThreadItem