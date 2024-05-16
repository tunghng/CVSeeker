
import ThreadMessageItem from "../ThreadMessageItem/ThreadMessageItem"

const ThreadMessageList = ({ threadMessages }) => {

    return (
        <div className="flex flex-col items-end">
            {threadMessages.data.map(message => (
                <ThreadMessageItem key={message.id} item={message} />
            ))}
        </div>
    )
}

export default ThreadMessageList