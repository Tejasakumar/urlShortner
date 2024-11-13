import React, { useState } from 'react';
import { Copy } from 'lucide-react';

const UrlShortener = () => {
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [copied, setCopied] = useState(false);
  const [error, setError] = useState('');

  const handleShorten = async () => {
    try {
      const response = await fetch(`/api/shorten?url=${encodeURIComponent(originalUrl)}`);
      const { shortUrl } = await response.json();
      setShortenedUrl(shortUrl);
      setError('');
    } catch {
      setError('Failed to shorten URL. Please try again later.');
      setShortenedUrl('');
    }
  };

  const handleCopy = () => {
    navigator.clipboard.writeText(shortenedUrl);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="bg-gray-900 p-6 rounded-lg shadow-md w-full lg:max-w-md">
      <h2 className="text-2xl font-semibold mb-4 text-white">Shorten Your URL</h2>
      <div className="mb-4">
        <div className="relative">
          <input
            type="text"
            placeholder="Enter URL to shorten"
            value={originalUrl}
            onChange={(e) => setOriginalUrl(e.target.value)}
            className="w-full py-2 px-4 rounded-md bg-gray-800 text-white border border-gray-600 focus:outline-none focus:ring-1 focus:ring-teal-500"
          />
          <button
            onClick={handleShorten}
            className="absolute right-2 top-1/2 -translate-y-1/2 bg-teal-500 hover:bg-teal-600 hover:scale-110 transition-transform duration-300 text-white px-4 py-2 rounded-md text-sm w-auto"
          >
            Shorten
          </button>
        </div>
      </div>
      {shortenedUrl && (
        <div className="flex items-center gap-2">
          <input
            readOnly
            value={shortenedUrl}
            className="flex-1 py-2 px-4 rounded-md bg-gray-800 text-white border border-gray-600"
          />
          <button
            onClick={handleCopy}
            className={`p-2 rounded-md animate-float ${
              copied ? 'bg-green-500 text-white' : 'bg-gray-700 text-white hover:bg-gray-600'
            }`}
          >
            {copied ? 'Copied!' : <Copy size={16} />}
          </button>
        </div>
      )}
      {error && (
        <div className="mt-4 py-2 px-3 rounded-md bg-red-900 text-sm text-red-400">
          {error}
        </div>
      )}
    </div>
  );
};

const App = () => {
  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-900">
      <UrlShortener />
    </div>
  );
};

export default App;