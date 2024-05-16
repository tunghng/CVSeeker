
import ReactMarkdown from 'react-markdown'

const ThreadMessageItem = ({ item }) => {

    return (
        <div className={`mt-4 px-4 py-3 flex rounded-xl
            ${item.role === 'user' ? 'bg-primary text-white ml-10' : 'bg-border'}
        `}>
            <ReactMarkdown>
                {item.content[0].text.value}
            </ReactMarkdown>
        </div>
    )
}

export default ThreadMessageItem