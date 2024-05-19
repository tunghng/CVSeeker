
import { Link } from "react-router-dom";

const PageButtons = ({ curr, query, level }) => {
    const total = 10
    curr = parseInt(curr, 10);

    const buttons = []
    let flag = true

    for (let i = 1; i <= total; i++) {
        if (i == 1 || i == total || i == curr || i == curr - 1 || i == curr - 2 || i == curr + 1 || i == curr + 2) {
            if (flag == false) {
                buttons.push(<div key={-i} className="text-text flex items-end ">...</div>)
                flag = true
            }
            buttons.push(
                <Link
                    className={`page-btn mx-2 ${(i == curr) && "page-btn-active"}`}
                    to={`/search?query=${query}&page=${i}&level=${level}`}
                    key={i}>
                    {i}
                </Link>
            )
        } else {
            flag = false
        }
    }

    return buttons
};

export default PageButtons;
