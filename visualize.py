import pandas as pd
import matplotlib.pyplot as plt
import time

df = pd.read_csv('bench.csv')

print(type(df))

insert_mask = df['operation'] == "Insert"
delete_mask = df['operation'] == "Delete"
find_mask = df['operation'] == "Find"
mem_mask = df['operation'] == "Memory"
insert_df = df[insert_mask]
delete_df = df[delete_mask]
find_df = df[find_mask]
mem_df = df[mem_mask]

#dfInsert := df

print(insert_df)

plt.rcParams['backend'] = "svg"

#print(df['# of elements'], "\n", df['duration (ns)'])

#plt.plot(df_interpolated['# of elements'], df_interpolated['duration (ns)'])
plt.plot(insert_df['# of elements'].to_numpy()[:,None], insert_df['duration (ns)'].to_numpy()[:,None])
plt.plot(delete_df['# of elements'].to_numpy()[:,None], delete_df['duration (ns)'].to_numpy()[:,None])
plt.plot(find_df['# of elements'].to_numpy()[:,None], find_df['duration (ns)'].to_numpy()[:,None])
plt.plot(mem_df['# of elements'].to_numpy()[:,None], mem_df['duration (ns)'].to_numpy()[:,None])
plt.xlabel('# of elements')
plt.ylabel('duration (ns)')
plt.title('Extrapolated Data')
plt.show()

#plt.plot([1,2,3])
#plt.show()
#time.sleep(10)
