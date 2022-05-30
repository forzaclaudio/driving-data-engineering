from asammdf import MDF, Signal
import mdfreader
import os
import pandas as pd
from functools import reduce
from pandas.api.types import is_numeric_dtype
from sklearn.preprocessing import OneHotEncoder


DATAPATH='/home/ec2-user/data/'
FILE='1CFD_VDDM_Lmbd1045_3.100_310_MS_log_SREC_log_shot0001_20190928_142405.dat'

def has_datadir(filename):
    """Check if filename contains path to datadir."""

    if DATAPATH in filename:
        return True
    return False

def add_datadir(filename):
    """Add path to datadir if not already in filename."""

    if not has_datadir(filename):
        return os.path.join(DATAPATH, FILE)
    else:
        return filename

def groups_in_file(mdf):
    """Return the number of groups in the MDF instance provided."""

    metadata = mdf.info()
    data_groups = 0
    for group in metadata.keys():
        is_dict = False
        try:
            metadata[group].keys()
            is_dict = True
        except:
            pass
        if is_dict:
            data_groups += 1
    print(data_groups, 'groups were found!')
    return data_groups

def display_metadata(mdf, signals_only=False):
    """Display the metadata of the MDF instance provided."""

    print('Groups:')
    for group in metadata.keys():
        print(group)
        is_dict = False
        try:
            metadata[group].keys()
            is_dict = True
        except:
            pass
        if is_dict:
            if signals_only:
                for i in range(1, metadata[group]['channels count']):
                    print(metadata[group]['channel '+str(i)])
            else:
                print(metadata[group])
        print('='*80)

def split_dataframe(full_df, debug=False):
    """Split dataframes in numeric and non numerical"""

    num_index = []
    non_num_index = []
    for i in range(full_df.shape[1]):
        if is_numeric_dtype(full_df.iloc[:,i]):
            if debug:
                print("Col", i, "is numeric")
            num_index.append(i)
        else:
            if debug:
                print("Col", i, "is non numeric")
            non_num_index.append(i)
    if debug:
        print("Numerical cols", num_index)
        print("Non numerical cols", non_num_index)
    numeric = full_df.iloc[:,num_index]
    non_num_index = full_df.iloc[:,non_num_index]
    return (numeric, non_num_index)

def extract_dataframes(data_groups, debug=False):
    """Extract numerical and non-numerical dataframes."""

    dfs = []
    non_numeric_dfs = []

    for g in range(data_groups):
        df = mdf.get_group(g)
        num, non_num = split_dataframe(df)
        if not num.empty:
            dfs.append(num)
        if not non_num.empty:
            non_numeric_dfs.append(non_num)
    return (dfs, non_numeric_dfs)

def hot_encode(input_df, debug=False):
    """Hot encode the given dataframe."""

    encoder = OneHotEncoder()
    encoded_df = pd.DataFrame(encoder.fit_transform(input_df).toarray(), index=input_df.index)

    column_names = {}
    final_names = []
    if debug:
        print("Input shape is:", input_df.shape[1])
    for col in input_df.columns:
        labels = []
        if debug:
            print("Generating column:", col)
        temp_labels = input_df[col].unique().tolist()
        for l in temp_labels:
            try:
                labels.append(l.decode('UTF-8'))
            except:
                labels.append(l)
        if debug:
            print("The following unique values were found:", labels)
        column_names[col] = labels
    if debug:
        print(column_names)
    for col in column_names.keys():
        for label in column_names[col]:
            final_names.append("{}_{}".format(col, label))
    encoded_df.columns = final_names
    return encoded_df

def hot_encode_dataframes(non_numeric_dfs, debug=False):
    """One-how encode the provided dataframes."""

    encoded_non_numeric_df = []
    for i, df in enumerate(non_numeric_dfs):
        if debug:
            print("Encoding dataframe:", i)
            print(df)
        temp = hot_encode(df, False)
        encoded_non_numeric_df.append(temp)
        if debug:
            print(temp)
    return encoded_non_numeric_df

def assemble_dataframe(dfs, debug=False):
    """Generate a global dataframe."""

    df_merged = reduce(lambda  left,right: pd.merge(left, right,
                       on=['timestamps'], how='outer'), dfs)
    if debug:
        print("Unsorted merge", df_merged.head(), df_merged.shape)
    df_merged_sorted = df_merged.sort_values('timestamps')
    if debug:
        print("Result of sorting", df_merged_sorted.head(), df_merged.shape)
    return df_merged_sorted

mdf = MDF(add_datadir(FILE))
data_groups = groups_in_file(mdf)
dfs, non_numeric_dfs = extract_dataframes(data_groups, debug=False)

print(len(dfs), "numerical dataframes found!")
print(len(non_numeric_dfs), "non-numerical dataframes found!")
if False:
    encoded_df = hot_encode_dataframes(non_numeric_dfs, debug=False)
    final_df = assemble_dataframe(encoded_df[50:-1], debug=False)
    final_df.to_csv('non_numerical_2_of_2.csv',index=True)

final_df = assemble_dataframe(dfs[100:-1], debug=False)
final_df.to_csv('numerical_3_of_3.csv',index=True)
