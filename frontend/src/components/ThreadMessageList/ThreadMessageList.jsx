
import ThreadMessageItem from "../ThreadMessageItem/ThreadMessageItem"

const ThreadMessageList = ({ threadMessages }) => {

    return (
        threadMessages.map(message => (
            <ThreadMessageItem key={message.id} item={message} />
        ))
    )
}

export default ThreadMessageList