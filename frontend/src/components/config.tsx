import { useEffect, useState } from 'react';

interface WebhookConfig {
  name: string;
  type: string;
  url: string;
  requestType: string;
  body: string;
  enabled: boolean;
}

interface WowConfig {
  basePath: string;
  enabled: boolean;
}

interface ConfigData {
  webhooks: WebhookConfig[];
  wow: WowConfig;
}

export default function Config(): JSX.Element {
  const [config, setConfig] = useState<ConfigData>({
    webhooks: [],
    wow: { basePath: '', enabled: false }
  });
  const [loading, setLoading] = useState(true);

  // Mock-Daten für die Demonstration - später durch echte API-Calls ersetzen
  useEffect(() => {
    // Simuliere das Laden der Konfiguration
    setTimeout(() => {
      setConfig({
        webhooks: [
          {
            name: "Discord Notification",
            type: "discord",
            url: "https://discord.com/api/webhooks/...",
            requestType: "POST",
            body: "Discord notification body",
            enabled: true
          },
          {
            name: "NTFY Notification", 
            type: "ntfy",
            url: "https://ntfy.sh/...",
            requestType: "POST",
            body: "NTFY notification body",
            enabled: false
          }
        ],
        wow: {
          basePath: "/Users/fabianfest/games/worldofwarcraft/",
          enabled: true
        }
      });
      setLoading(false);
    }, 500);
  }, []);

  const toggleWebhook = (index: number) => {
    setConfig(prev => ({
      ...prev,
      webhooks: prev.webhooks.map((webhook, i) => 
        i === index ? { ...webhook, enabled: !webhook.enabled } : webhook
      )
    }));
  };

  const toggleWow = () => {
    setConfig(prev => ({
      ...prev,
      wow: { ...prev.wow, enabled: !prev.wow.enabled }
    }));
  };

  const saveConfig = () => {
    // TODO: Implementiere das Speichern der Konfiguration über die Go-API
    console.log('Saving config:', config);
    // Hier würde normalerweise ein API-Call zur Go-Backend erfolgen
  };

  if (loading) {
    return (
      <div className="p-6">
        <div className="animate-pulse">
          <div className="h-4 bg-gray-200 rounded w-1/4 mb-4"></div>
          <div className="space-y-3">
            <div className="h-4 bg-gray-200 rounded"></div>
            <div className="h-4 bg-gray-200 rounded w-5/6"></div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <div className="bg-white rounded-lg shadow-md">
        <div className="px-6 py-4 border-b border-gray-200">
          <h2 className="text-xl font-semibold text-gray-800">Konfiguration</h2>
          <p className="text-sm text-gray-600 mt-1">
            Verwalten Sie Ihre Webhook-Benachrichtigungen und WoW-Einstellungen
          </p>
        </div>

        <div className="p-6 space-y-6">
          {/* Webhooks Section */}
          <div>
            <h3 className="text-lg font-medium text-gray-900 mb-4">
              Webhook-Benachrichtigungen
            </h3>
            <div className="space-y-3">
              {config.webhooks.map((webhook, index) => (
                <div 
                  key={index}
                  className={`p-4 border rounded-lg transition-colors ${
                    webhook.enabled 
                      ? 'border-green-200 bg-green-50' 
                      : 'border-gray-200 bg-gray-50'
                  }`}
                >
                  <div className="flex items-center justify-between">
                    <div className="flex-1">
                      <div className="flex items-center space-x-3">
                        <label className="flex items-center cursor-pointer">
                          <input
                            type="checkbox"
                            checked={webhook.enabled}
                            onChange={() => toggleWebhook(index)}
                            className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                          />
                          <span className="ml-2 text-sm font-medium text-gray-900">
                            {webhook.name}
                          </span>
                        </label>
                        <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                          webhook.type === 'discord' 
                            ? 'bg-indigo-100 text-indigo-800'
                            : webhook.type === 'ntfy'
                            ? 'bg-purple-100 text-purple-800'
                            : 'bg-gray-100 text-gray-800'
                        }`}>
                          {webhook.type}
                        </span>
                      </div>
                      <div className="mt-2 text-sm text-gray-600">
                        <div className="font-mono text-xs bg-gray-100 px-2 py-1 rounded truncate">
                          {webhook.url}
                        </div>
                      </div>
                    </div>
                    <div className="ml-4">
                      <span className={`inline-flex items-center px-2 py-1 rounded-full text-xs font-medium ${
                        webhook.enabled 
                          ? 'bg-green-100 text-green-800' 
                          : 'bg-red-100 text-red-800'
                      }`}>
                        {webhook.enabled ? 'Aktiv' : 'Inaktiv'}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* WoW Configuration Section */}
          <div>
            <h3 className="text-lg font-medium text-gray-900 mb-4">
              World of Warcraft
            </h3>
            <div className={`p-4 border rounded-lg transition-colors ${
              config.wow.enabled 
                ? 'border-green-200 bg-green-50' 
                : 'border-gray-200 bg-gray-50'
            }`}>
              <div className="flex items-center justify-between">
                <div className="flex-1">
                  <div className="flex items-center space-x-3">
                    <label className="flex items-center cursor-pointer">
                      <input
                        type="checkbox"
                        checked={config.wow.enabled}
                        onChange={toggleWow}
                        className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                      />
                      <span className="ml-2 text-sm font-medium text-gray-900">
                        WoW Queue Monitoring
                      </span>
                    </label>
                  </div>
                  <div className="mt-2 text-sm text-gray-600">
                    <div className="font-mono text-xs bg-gray-100 px-2 py-1 rounded">
                      Pfad: {config.wow.basePath}
                    </div>
                  </div>
                </div>
                <div className="ml-4">
                  <span className={`inline-flex items-center px-2 py-1 rounded-full text-xs font-medium ${
                    config.wow.enabled 
                      ? 'bg-green-100 text-green-800' 
                      : 'bg-red-100 text-red-800'
                  }`}>
                    {config.wow.enabled ? 'Aktiv' : 'Inaktiv'}
                  </span>
                </div>
              </div>
            </div>
          </div>

          {/* Save Button */}
          <div className="pt-4 border-t border-gray-200">
            <div className="flex justify-end">
              <button
                onClick={saveConfig}
                className="px-4 py-2 bg-blue-600 text-white text-sm font-medium rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
              >
                Konfiguration speichern
              </button>
            </div>
          </div>

          {/* Summary */}
          <div className="bg-blue-50 border border-blue-200 rounded-lg p-4">
            <h4 className="text-sm font-medium text-blue-900 mb-2">Zusammenfassung</h4>
            <div className="text-sm text-blue-800">
              <div className="flex justify-between">
                <span>Aktive Webhooks:</span>
                <span className="font-medium">
                  {config.webhooks.filter(w => w.enabled).length} von {config.webhooks.length}
                </span>
              </div>
              <div className="flex justify-between">
                <span>WoW Monitoring:</span>
                <span className="font-medium">
                  {config.wow.enabled ? 'Aktiviert' : 'Deaktiviert'}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
