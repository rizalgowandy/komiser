import { Tag } from '../hooks/useInventory/types/useInventoryTypes';

type InventoryTableTagsProps = {
  tags: [] | Tag[] | null;
  setQuery: (query: string) => void;
  id: string;
  bulkItems: [] | string[];
};

function InventoryTableTags({
  tags,
  setQuery,
  id,
  bulkItems
}: InventoryTableTagsProps) {
  return (
    <>
      {tags && tags.length > 0 && (
        <div className="group relative">
          <div className="relative flex items-center gap-1 py-4 px-6">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="none"
              viewBox="0 0 24 24"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
              ></path>
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeWidth="1.5"
                d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
              ></path>
            </svg>
            <span
              className={`text-black-900" absolute top-3 left-[2.375rem] flex h-4 w-4 items-center justify-center rounded-full bg-white text-[10px] font-bold ${
                bulkItems && bulkItems.find(currentId => currentId === id)
                  ? 'bg-komiser-120'
                  : ''
              }`}
            >
              {tags.length}
            </span>
          </div>
          <div className="absolute right-6 top-11 z-10 hidden flex-col gap-2 rounded-lg bg-black-900 py-3 px-4 shadow-lg group-hover:flex">
            {tags.map((tag, index) => (
              <div
                key={index}
                className="-mx-4 flex items-center gap-2 border-t border-white/20 px-4 pt-2 text-xs text-black-200 first:border-none first:pt-0"
              >
                <div className="flex items-center gap-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="14"
                    height="14"
                    fill="none"
                    viewBox="0 0 24 24"
                    className="text-black-400"
                  >
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="1.5"
                      d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
                    ></path>
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeWidth="1.5"
                      d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                    ></path>
                  </svg>
                  <span
                    onClick={e => {
                      setQuery(tag.key);
                    }}
                    className="cursor-pointer hover:text-komiser-500"
                  >
                    {tag.key}:
                  </span>
                </div>
                <span
                  onClick={() => setQuery(tag.value)}
                  className="cursor-pointer font-medium hover:text-komiser-500"
                >
                  {tag.value}
                </span>
              </div>
            ))}
          </div>
        </div>
      )}
    </>
  );
}

export default InventoryTableTags;
