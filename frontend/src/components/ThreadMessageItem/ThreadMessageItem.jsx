
import ReactMarkdown from 'react-markdown'

const ThreadMessageItem = ({ item }) => {
    if (item && item.content[0].text.value.startsWith('You will use these information')) {
        return null;
    }

    return (
        <div className={`mt-5 px-4 py-2.5 rounded-xl
            ${item.role === 'user' ? 'bg-primary text-white ml-10 self-end' : 'bg-border self-start'}
        `}>
            <ReactMarkdown>
                {item.content[0].text.value}
            </ReactMarkdown>
        </div>
    )
}

export default ThreadMessageItem
